package contract_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"
	"time"

	ofauth "github.com/ofauth-org/onlyfans-sdk-go"
)

type endpointSpec struct {
	Method string `json:"method"`
	Route  string `json:"route"`
}

func loadDotEnvLocal() {
	// Optional convenience: load `.env.e2e.local` from current or parent dirs if present,
	// without overriding already-set environment variables.
	tryPaths := []string{
		".env.e2e.local",
		"../.env.e2e.local",
		"../../.env.e2e.local",
		"../../../.env.e2e.local",
		"../../../../.env.e2e.local",
	}

	for _, p := range tryPaths {
		raw, err := os.ReadFile(p)
		if err != nil {
			continue
		}

		lines := strings.Split(string(raw), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			eq := strings.IndexByte(line, '=')
			if eq <= 0 {
				continue
			}
			key := strings.TrimSpace(line[:eq])
			val := strings.TrimSpace(line[eq+1:])
			if key == "" {
				continue
			}
			if len(val) >= 2 {
				if (val[0] == '"' && val[len(val)-1] == '"') || (val[0] == '\'' && val[len(val)-1] == '\'') {
					val = val[1 : len(val)-1]
				}
			}
			if _, ok := os.LookupEnv(key); !ok {
				_ = os.Setenv(key, val)
			}
		}
		return
	}
}

func splitCSV(v string) []string {
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	return out
}

func shouldRunSuite(name string) bool {
	suites := splitCSV(os.Getenv("E2E_SDK_SUITES"))
	if len(suites) == 0 {
		return true // default: run getters
	}
	for _, s := range suites {
		if s == name {
			return true
		}
	}
	return false
}

func defaultQuery(route string, fixtures map[string]string) url.Values {
	if route == "/users/search" {
		return url.Values{"query": []string{"test"}}
	}

	if route == "/users/list" {
		if userId := fixtures["userId"]; userId != "" {
			return url.Values{"userIds": []string{userId}}
		}
		return nil
	}

	if strings.HasSuffix(route, "/chart") {
		return url.Values{
			"startDate": []string{"2024-01-01"},
			"endDate":   []string{"2024-01-31"},
			"withTotal": []string{"true"},
		}
	}

	if strings.HasSuffix(route, "/top") {
		return url.Values{
			"startDate": []string{"2024-01-01"},
			"endDate":   []string{"2024-01-31"},
			"limit":     []string{"1"},
			"offset":    []string{"0"},
		}
	}

	if route == "/posts" ||
		route == "/subscribers" ||
		route == "/subscriptions" ||
		route == "/chats" ||
		route == "/mass-messages" ||
		route == "/analytics/mass-messages/purchased" ||
		route == "/earnings/transactions" ||
		route == "/earnings/chargebacks" ||
		route == "/promotions" ||
		route == "/promotions/bundles" ||
		route == "/promotions/tracking-links" ||
		route == "/promotions/trial-links" ||
		route == "/users/blocked" ||
		route == "/users/restrict" ||
		route == "/users/lists" ||
		route == "/vault/media" ||
		route == "/vault/lists" ||
		strings.HasSuffix(route, "/notifications") ||
		strings.HasSuffix(route, "/release-forms") ||
		strings.HasSuffix(route, "/tagged-friend-users") ||
		strings.HasSuffix(route, "/users") ||
		strings.HasSuffix(route, "/buyers") {
		return url.Values{"limit": []string{"1"}}
	}

	return nil
}

func routeParamNames(routeTemplate string) []string {
	re := regexp.MustCompile(`:([A-Za-z0-9_]+)`)
	matches := re.FindAllStringSubmatch(routeTemplate, -1)
	out := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) < 2 {
			continue
		}
		out = append(out, m[1])
	}
	return out
}

func fillRouteParams(routeTemplate string, params map[string]string) string {
	re := regexp.MustCompile(`:([A-Za-z0-9_]+)`)
	return re.ReplaceAllStringFunc(routeTemplate, func(m string) string {
		sub := re.FindStringSubmatch(m)
		if len(sub) < 2 {
			return m
		}
		if v, ok := params[sub[1]]; ok {
			return v
		}
		return m
	})
}

func fixtureKeyForRouteParam(route string, paramName string) string {
	if paramName == "userId" && strings.HasPrefix(route, "/chats/") {
		return "chatUserId"
	}
	if paramName == "listId" && strings.HasPrefix(route, "/vault/lists/") {
		return "vaultListId"
	}
	return paramName
}

func asMap(v any) (map[string]any, bool) {
	m, ok := v.(map[string]any)
	return m, ok
}

func asSlice(v any) ([]any, bool) {
	s, ok := v.([]any)
	return s, ok
}

func selfID(v any) (string, bool) {
	obj, ok := asMap(v)
	if !ok {
		return "", false
	}
	if id, ok := obj["id"]; ok && id != nil {
		return fmt.Sprintf("%v", id), true
	}
	if user, ok := obj["user"]; ok {
		if m, ok := asMap(user); ok {
			if id, ok := m["id"]; ok && id != nil {
				return fmt.Sprintf("%v", id), true
			}
		}
	}
	if data, ok := obj["data"]; ok {
		if m, ok := asMap(data); ok {
			if id, ok := m["id"]; ok && id != nil {
				return fmt.Sprintf("%v", id), true
			}
		}
	}
	return "", false
}

func firstListItemID(v any) (string, bool) {
	obj, ok := asMap(v)
	if !ok {
		return "", false
	}
	listAny := obj["list"]
	if listAny == nil {
		listAny = obj["items"]
	}
	list, ok := asSlice(listAny)
	if !ok || len(list) == 0 {
		return "", false
	}
	first, ok := asMap(list[0])
	if !ok {
		return "", false
	}
	for _, key := range []string{
		"id",
		"postId",
		"listId",
		"trackingLinkId",
		"trialLinkId",
		"promotionId",
		"bundleId",
		"subscriptionId",
		"massMessageId",
	} {
		if val, ok := first[key]; ok && val != nil {
			return fmt.Sprintf("%v", val), true
		}
	}
	return "", false
}

func chatUserIDFromChats(v any) (string, bool) {
	obj, ok := asMap(v)
	if !ok {
		return "", false
	}
	listAny := obj["list"]
	if listAny == nil {
		listAny = obj["items"]
	}
	list, ok := asSlice(listAny)
	if !ok || len(list) == 0 {
		return "", false
	}
	first, ok := asMap(list[0])
	if !ok {
		return "", false
	}
	if withUser, ok := first["withUser"]; ok {
		if m, ok := asMap(withUser); ok {
			if id, ok := m["id"]; ok && id != nil {
				return fmt.Sprintf("%v", id), true
			}
		}
	}
	if user, ok := first["user"]; ok {
		if m, ok := asMap(user); ok {
			if id, ok := m["id"]; ok && id != nil {
				return fmt.Sprintf("%v", id), true
			}
		}
	}
	for _, key := range []string{"userId", "withUserId", "id"} {
		if val, ok := first[key]; ok && val != nil {
			return fmt.Sprintf("%v", val), true
		}
	}
	return "", false
}

func sleepUntilNext(minDelayMs int, lastStart time.Time) time.Time {
	if minDelayMs <= 0 {
		return time.Now()
	}
	next := lastStart.Add(time.Duration(minDelayMs) * time.Millisecond)
	now := time.Now()
	if now.Before(next) {
		time.Sleep(next.Sub(now))
	}
	return time.Now()
}

func TestAccessGetContractGeneratedSDK(t *testing.T) {
	if !shouldRunSuite("access-get") {
		t.Skip("suite disabled via E2E_SDK_SUITES")
	}

	loadDotEnvLocal()

	raw, err := os.ReadFile("accessEndpoints.manifest.json")
	if err != nil {
		t.Fatalf("read manifest: %v", err)
	}

	var manifest []endpointSpec
	if err := json.Unmarshal(raw, &manifest); err != nil {
		t.Fatalf("parse manifest: %v", err)
	}

	var getEndpoints []endpointSpec
	for _, e := range manifest {
		if strings.ToLower(e.Method) == "get" {
			getEndpoints = append(getEndpoints, e)
		}
	}
	if len(getEndpoints) < 20 {
		t.Fatalf("unexpected manifest get endpoint count: %d", len(getEndpoints))
	}

	baseURL := strings.TrimSpace(os.Getenv("E2E_ACCESS_BASE_URL"))
	required := strings.EqualFold(os.Getenv("E2E_CONTRACT_REQUIRED"), "true") || os.Getenv("E2E_CONTRACT_REQUIRED") == "1"
	strict := strings.EqualFold(os.Getenv("E2E_CONTRACT_STRICT"), "true") || os.Getenv("E2E_CONTRACT_STRICT") == "1"
	runLive := strings.EqualFold(os.Getenv("E2E_RUN_LIVE"), "true") || os.Getenv("E2E_RUN_LIVE") == "1"

	type ctxSpec struct {
		Name       string
		APIKey     string
		Connection string
		MinDelayMs int
	}

	contexts := []ctxSpec{
		{
			Name:       "sandbox-creator",
			APIKey:     os.Getenv("E2E_SANDBOX_API_KEY_CREATOR"),
			Connection: os.Getenv("E2E_SANDBOX_CONNECTION_ID_CREATOR"),
			MinDelayMs: 0,
		},
		{
			Name:       "sandbox-fan",
			APIKey:     os.Getenv("E2E_SANDBOX_API_KEY_FAN"),
			Connection: os.Getenv("E2E_SANDBOX_CONNECTION_ID_FAN"),
			MinDelayMs: 0,
		},
	}
	if runLive {
		contexts = append(contexts,
			ctxSpec{
				Name:       "live-creator",
				APIKey:     os.Getenv("E2E_LIVE_API_KEY_CREATOR"),
				Connection: os.Getenv("E2E_LIVE_CONNECTION_ID_CREATOR"),
				MinDelayMs: 2000,
			},
			ctxSpec{
				Name:       "live-fan",
				APIKey:     os.Getenv("E2E_LIVE_API_KEY_FAN"),
				Connection: os.Getenv("E2E_LIVE_CONNECTION_ID_FAN"),
				MinDelayMs: 2000,
			},
		)
	}

	remaining := make(map[string]bool, len(getEndpoints))
	for _, e := range getEndpoints {
		remaining[e.Route] = true
	}

	ranAny := false

	for _, ctxCfg := range contexts {
		if baseURL == "" || ctxCfg.APIKey == "" || ctxCfg.Connection == "" {
			if required {
				t.Fatalf("missing env for context %s; need E2E_ACCESS_BASE_URL + api key + connection id", ctxCfg.Name)
			}
			continue
		}

		ranAny = true

		httpClient := &http.Client{Timeout: 60 * time.Second}
		client := ofauth.NewClient(ctxCfg.APIKey).
			WithBaseURL(baseURL).
			WithConnectionID(ctxCfg.Connection).
			WithHTTPClient(httpClient)

		fixtures := map[string]string{}
		var lastStart time.Time

		rateLimitedGet := func(path string, q url.Values) (any, []byte, error) {
			lastStart = sleepUntilNext(ctxCfg.MinDelayMs, lastStart)
			body, err := client.Request(context.Background(), "GET", path, &ofauth.RequestOptions{Query: q, ConnectionID: ctxCfg.Connection})
			if err != nil {
				return nil, nil, err
			}
			var parsed any
			if err := json.Unmarshal(body, &parsed); err != nil {
				return nil, body, err
			}
			return parsed, body, nil
		}

		// Preflight required-success call.
		selfParsed, _, err := rateLimitedGet("/v2/access/self", nil)
		if err != nil {
			t.Fatalf("[%s] preflight /self failed: %v", ctxCfg.Name, err)
		}
		if id, ok := selfID(selfParsed); ok {
			fixtures["userId"] = id
		}

		ensureFixture := func(name string) (string, bool, error) {
			if v, ok := fixtures[name]; ok && v != "" {
				return v, true, nil
			}

			var (
				parsed any
				err    error
			)

			switch name {
			case "userId":
				if v, ok := fixtures["userId"]; ok && v != "" {
					return v, true, nil
				}
				return "", false, nil
			case "chatUserId":
				parsed, _, err = rateLimitedGet("/v2/access/chats", url.Values{"limit": []string{"1"}})
				if err != nil {
					return "", false, err
				}
				if id, ok := chatUserIDFromChats(parsed); ok {
					fixtures[name] = id
					return id, true, nil
				}
				return "", false, nil
			case "postId":
				parsed, _, err = rateLimitedGet("/v2/access/posts", url.Values{"limit": []string{"1"}})
			case "listId":
				parsed, _, err = rateLimitedGet("/v2/access/users/lists", url.Values{"limit": []string{"1"}})
			case "vaultListId":
				parsed, _, err = rateLimitedGet("/v2/access/vault/lists", url.Values{"limit": []string{"1"}})
			case "subscriptionId":
				parsed, _, err = rateLimitedGet("/v2/access/subscriptions", url.Values{"limit": []string{"1"}})
			case "massMessageId":
				parsed, _, err = rateLimitedGet("/v2/access/mass-messages", url.Values{"limit": []string{"1"}})
			case "promotionId":
				parsed, _, err = rateLimitedGet("/v2/access/promotions", url.Values{"limit": []string{"1"}})
			case "bundleId":
				parsed, _, err = rateLimitedGet("/v2/access/promotions/bundles", url.Values{"limit": []string{"1"}})
			case "trackingLinkId":
				parsed, _, err = rateLimitedGet("/v2/access/promotions/tracking-links", url.Values{"limit": []string{"1"}})
			case "trialLinkId":
				parsed, _, err = rateLimitedGet("/v2/access/promotions/trial-links", url.Values{"limit": []string{"1"}})
			default:
				return "", false, nil
			}

			if err != nil {
				return "", false, err
			}

			if id, ok := firstListItemID(parsed); ok {
				fixtures[name] = id
				return id, true, nil
			}
			return "", false, nil
		}

		for _, ep := range getEndpoints {
			paramNames := routeParamNames(ep.Route)
			params := map[string]string{}
			canRun := true
			for _, p := range paramNames {
				fixtureKey := fixtureKeyForRouteParam(ep.Route, p)
				v, ok, err := ensureFixture(fixtureKey)
				if err != nil {
					t.Fatalf("[%s] fixture discovery failed (%s for route=%s): %v", ctxCfg.Name, fixtureKey, ep.Route, err)
				}
				if !ok || v == "" {
					canRun = false
					break
				}
				params[p] = v
			}
			if !canRun {
				continue
			}

			path := "/v2/access" + fillRouteParams(ep.Route, params)
			q := defaultQuery(ep.Route, fixtures)

			parsed, body, err := rateLimitedGet(path, q)
			if err != nil {
				t.Fatalf("[%s] GET %s failed: %v", ctxCfg.Name, path, err)
			}

			if _, ok := parsed.(string); ok {
				if regexp.MustCompile(`"url"\s*:\s*"{3,}`).Find(body) != nil {
					t.Fatalf("[%s] raw string contains malformed JSON (e.g. url:\"\"\"\") route=%s", ctxCfg.Name, ep.Route)
				}
				t.Fatalf("[%s] raw string returned (expected parsed JSON) route=%s", ctxCfg.Name, ep.Route)
			}

			delete(remaining, ep.Route)
		}
	}

	if !ranAny {
		t.Skip("no e2e credentials provided (set E2E_ACCESS_BASE_URL + connection/api keys or E2E_CONTRACT_REQUIRED=1)")
	}

	if strict && len(remaining) > 0 {
		missing := make([]string, 0, len(remaining))
		for r := range remaining {
			missing = append(missing, r)
		}
		sort.Strings(missing)
		t.Fatalf("missing happy-path coverage for %d GET access endpoints:\n- %s", len(missing), strings.Join(missing, "\n- "))
	}
}
