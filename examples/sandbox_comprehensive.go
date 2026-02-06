package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	ofauth "github.com/ofauth-org/ofauth-go"
)

var (
	apiKey       = os.Getenv("OFAUTH_API_KEY")
	connectionID = os.Getenv("OFAUTH_CONNECTION_ID")
	baseURL      = getEnvOrDefault("OFAUTH_BASE_URL", "https://api.ofauth.com")
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func init() {
	if apiKey == "" || connectionID == "" {
		fmt.Println("Missing required environment variables:")
		fmt.Println("  OFAUTH_API_KEY - Your OFAuth API key")
		fmt.Println("  OFAUTH_CONNECTION_ID - Your connection ID")
		fmt.Println("  OFAUTH_BASE_URL - (optional) API base URL")
		os.Exit(1)
	}
}

type testResult struct {
	passed int
	failed int
	errors []struct {
		test  string
		error string
	}
}

var results = &testResult{}

func success(name string, data interface{}) {
	results.passed++
	fmt.Printf("  ✓ %s\n", name)
}

func fail(name string, err error) {
	results.failed++
	results.errors = append(results.errors, struct {
		test  string
		error string
	}{name, err.Error()})
	fmt.Printf("  ✗ %s: %s\n", name, err.Error())
}

func testAccountAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n📋 Account API")
	fmt.Println("----------------------------------------")

	if data, err := client.Whoami(ctx); err != nil {
		fail("Whoami", err)
	} else {
		success("Whoami", data)
		fmt.Printf("    → id: %s\n", data.Id)
		fmt.Printf("    → name: %s\n", data.Name)
	}

	if data, err := client.ListAccountConnections(ctx, nil); err != nil {
		fail("ListAccountConnections", err)
	} else {
		success("ListAccountConnections", data)
		fmt.Printf("    → list: [%d items]\n", len(data.List))
		fmt.Printf("    → hasMore: %v\n", data.HasMore)
	}

	if data, err := client.GetAccountConnectionsSettings(ctx, connectionID); err != nil {
		fail("GetAccountConnectionsSettings", err)
	} else {
		success("GetAccountConnectionsSettings", data)
	}

	if data, err := client.GetAccountSettings(ctx); err != nil {
		fail("GetAccountSettings", err)
	} else {
		success("GetAccountSettings", data)
	}
}

func testSelfAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n👤 Self API")
	fmt.Println("----------------------------------------")

	if data, err := client.GetAccessSelf(ctx); err != nil {
		fail("GetAccessSelf", err)
	} else {
		success("GetAccessSelf", data)
		if id, ok := (*data)["id"]; ok {
			fmt.Printf("    → id: %v\n", id)
		}
		if username, ok := (*data)["username"]; ok {
			fmt.Printf("    → username: %v\n", username)
		}
	}

	opts := &ofauth.RequestOptions{Query: url.Values{"limit": {"10"}}}
	if data, err := client.ListAccessSelfNotifications(ctx, opts); err != nil {
		fail("ListAccessSelfNotifications", err)
	} else {
		success("ListAccessSelfNotifications", data)
	}

	if data, err := client.ListAccessSelfReleaseForms(ctx, opts); err != nil {
		fail("ListAccessSelfReleaseForms", err)
	} else {
		success("ListAccessSelfReleaseForms", data)
	}

	if data, err := client.ListAccessSelfTaggedFriendUsers(ctx, opts); err != nil {
		fail("ListAccessSelfTaggedFriendUsers", err)
	} else {
		success("ListAccessSelfTaggedFriendUsers", data)
	}
}

func testPostsAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n📝 Posts API")
	fmt.Println("----------------------------------------")

	opts := &ofauth.RequestOptions{Query: url.Values{"limit": {"10"}}}
	if data, err := client.ListAccessPosts(ctx, opts); err != nil {
		fail("ListAccessPosts", err)
	} else {
		success("ListAccessPosts", data)
		fmt.Printf("    → list: [%d items]\n", len(data.List))
		fmt.Printf("    → hasMore: %v\n", data.HasMore)

		if len(data.List) > 0 {
			postId := fmt.Sprintf("%v", data.List[0].Id)
			if post, err := client.GetAccessPosts(ctx, postId); err != nil {
				fail(fmt.Sprintf("GetAccessPosts (id=%s)", postId), err)
			} else {
				success(fmt.Sprintf("GetAccessPosts (id=%s)", postId), post)
			}
		}
	}
}

func testMessagesAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n💬 Messages API")
	fmt.Println("----------------------------------------")

	opts := &ofauth.RequestOptions{Query: url.Values{"limit": {"10"}}}
	if data, err := client.ListAccessChats(ctx, opts); err != nil {
		fail("ListAccessChats", err)
	} else {
		success("ListAccessChats", data)
	}

	if data, err := client.ListAccessMassMessages(ctx, opts); err != nil {
		fail("ListAccessMassMessages", err)
	} else {
		success("ListAccessMassMessages", data)
	}
}

func testSubscribersAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n👥 Subscribers API")
	fmt.Println("----------------------------------------")

	opts := &ofauth.RequestOptions{Query: url.Values{"type": {"all"}, "limit": {"10"}}}
	if data, err := client.ListAccessSubscribers(ctx, opts); err != nil {
		fail("ListAccessSubscribers (all)", err)
	} else {
		success("ListAccessSubscribers (all)", data)
	}

	opts.Query.Set("type", "active")
	if data, err := client.ListAccessSubscribers(ctx, opts); err != nil {
		fail("ListAccessSubscribers (active)", err)
	} else {
		success("ListAccessSubscribers (active)", data)
	}
}

func testSubscriptionsAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n📊 Subscriptions API")
	fmt.Println("----------------------------------------")

	if data, err := client.GetAccessSubscriptionsCount(ctx); err != nil {
		fail("GetAccessSubscriptionsCount", err)
	} else {
		success("GetAccessSubscriptionsCount", data)
	}

	opts := &ofauth.RequestOptions{Query: url.Values{"limit": {"10"}}}
	if data, err := client.ListAccessSubscriptions(ctx, opts); err != nil {
		fail("ListAccessSubscriptions", err)
	} else {
		success("ListAccessSubscriptions", data)
	}
}

func testUsersAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n🔍 Users API")
	fmt.Println("----------------------------------------")

	opts := &ofauth.RequestOptions{Query: url.Values{"limit": {"10"}}}
	if data, err := client.GetAccessUsersBlocked(ctx, opts); err != nil {
		fail("GetAccessUsersBlocked", err)
	} else {
		success("GetAccessUsersBlocked", data)
	}

	if data, err := client.GetAccessUsersRestrict(ctx, opts); err != nil {
		fail("GetAccessUsersRestrict", err)
	} else {
		success("GetAccessUsersRestrict", data)
	}

	opts.Query.Set("query", "test")
	if data, err := client.AccessUsersSearch(ctx, opts); err != nil {
		fail("AccessUsersSearch", err)
	} else {
		success("AccessUsersSearch", data)
	}
}

func testUserListsAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n📁 User Lists API")
	fmt.Println("----------------------------------------")

	opts := &ofauth.RequestOptions{Query: url.Values{"limit": {"10"}}}
	if data, err := client.ListAccessUsersLists(ctx, opts); err != nil {
		fail("ListAccessUsersLists", err)
	} else {
		success("ListAccessUsersLists", data)
	}
}

func testVaultAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n🗄️ Vault API")
	fmt.Println("----------------------------------------")

	opts := &ofauth.RequestOptions{Query: url.Values{"limit": {"10"}}}
	if data, err := client.ListAccessVaultMedia(ctx, opts); err != nil {
		fail("ListAccessVaultMedia", err)
	} else {
		success("ListAccessVaultMedia", data)
	}

	if data, err := client.ListAccessVaultLists(ctx, opts); err != nil {
		fail("ListAccessVaultLists", err)
	} else {
		success("ListAccessVaultLists", data)
	}
}

func testEarningsAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n💰 Earnings API")
	fmt.Println("----------------------------------------")

	now := time.Now()
	thirtyDaysAgo := now.AddDate(0, 0, -30)
	startDate := thirtyDaysAgo.Format("2006-01-02")
	endDate := now.Format("2006-01-02")

	opts := &ofauth.RequestOptions{Query: url.Values{"startDate": {startDate}, "endDate": {endDate}}}
	if data, err := client.GetAccessEarningsChart(ctx, opts); err != nil {
		fail("GetAccessEarningsChart", err)
	} else {
		success("GetAccessEarningsChart", data)
	}

	opts.Query.Del("endDate")
	if data, err := client.ListAccessEarningsTransactions(ctx, opts); err != nil {
		fail("ListAccessEarningsTransactions", err)
	} else {
		success("ListAccessEarningsTransactions", data)
	}

	opts.Query.Set("endDate", endDate)
	opts.Query.Set("limit", "10")
	if data, err := client.ListAccessEarningsChargebacks(ctx, opts); err != nil {
		fail("ListAccessEarningsChargebacks", err)
	} else {
		success("ListAccessEarningsChargebacks", data)
	}
}

func testAnalyticsAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n📈 Analytics API")
	fmt.Println("----------------------------------------")

	now := time.Now()
	thirtyDaysAgo := now.AddDate(0, 0, -30)
	startDate := thirtyDaysAgo.Format("2006-01-02")
	endDate := now.Format("2006-01-02")

	opts := &ofauth.RequestOptions{Query: url.Values{"startDate": {startDate}, "endDate": {endDate}}}
	if data, err := client.GetAccessAnalyticsPostsChart(ctx, opts); err != nil {
		fail("GetAccessAnalyticsPostsChart", err)
	} else {
		success("GetAccessAnalyticsPostsChart", data)
	}

	opts.Query.Set("limit", "10")
	if data, err := client.GetAccessAnalyticsPostsTop(ctx, opts); err != nil {
		fail("GetAccessAnalyticsPostsTop", err)
	} else {
		success("GetAccessAnalyticsPostsTop", data)
	}

	if data, err := client.GetAccessAnalyticsStoriesChart(ctx, opts); err != nil {
		fail("GetAccessAnalyticsStoriesChart", err)
	} else {
		success("GetAccessAnalyticsStoriesChart", data)
	}
}

func testPromotionsAPI(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n🎁 Promotions API")
	fmt.Println("----------------------------------------")

	opts := &ofauth.RequestOptions{Query: url.Values{"limit": {"10"}}}
	if data, err := client.ListAccessPromotions(ctx, opts); err != nil {
		fail("ListAccessPromotions", err)
	} else {
		success("ListAccessPromotions", data)
	}

	if data, err := client.ListAccessPromotionsTrackingLinks(ctx, opts); err != nil {
		fail("ListAccessPromotionsTrackingLinks", err)
	} else {
		success("ListAccessPromotionsTrackingLinks", data)
	}

	if data, err := client.ListAccessPromotionsTrialLinks(ctx, opts); err != nil {
		fail("ListAccessPromotionsTrialLinks", err)
	} else {
		success("ListAccessPromotionsTrialLinks", data)
	}

	if data, err := client.ListAccessPromotionsBundles(ctx, opts); err != nil {
		fail("ListAccessPromotionsBundles", err)
	} else {
		success("ListAccessPromotionsBundles", data)
	}
}

func testMutationAPIs(client *ofauth.Client, ctx context.Context) {
	fmt.Println("\n✏️ Mutation APIs (CRUD Operations)")
	fmt.Println("----------------------------------------")

	var createdVaultListId string

	created, err := client.CreateAccessVaultLists(ctx, struct {
		Name string `json:"name"`
	}{Name: fmt.Sprintf("SDK Test List %d", time.Now().Unix())})
	if err != nil {
		fail("CreateAccessVaultLists", err)
	} else {
		createdVaultListId = fmt.Sprintf("%v", created.Id)
		success("CreateAccessVaultLists", created)
		fmt.Printf("    → id: %s\n", createdVaultListId)
	}

	if createdVaultListId != "" {
		updated, err := client.UpdateAccessVaultLists(ctx, createdVaultListId, struct {
			Name       string `json:"name"`
			ClearMedia *bool  `json:"clearMedia,omitempty"`
		}{Name: fmt.Sprintf("SDK Test List Updated %d", time.Now().Unix())})
		if err != nil {
			fail("UpdateAccessVaultLists", err)
		} else {
			success("UpdateAccessVaultLists", updated)
		}

		if _, err := client.DeleteAccessVaultLists(ctx, createdVaultListId); err != nil {
			fail("DeleteAccessVaultLists", err)
		} else {
			success("DeleteAccessVaultLists", nil)
		}
	}

	testUserId := "12345"
	if _, err := client.CreateAccessUsersRestrict(ctx, testUserId); err != nil {
		fail("CreateAccessUsersRestrict", err)
	} else {
		success("CreateAccessUsersRestrict", nil)
	}

	if _, err := client.DeleteAccessUsersRestrict(ctx, testUserId); err != nil {
		fail("DeleteAccessUsersRestrict", err)
	} else {
		success("DeleteAccessUsersRestrict", nil)
	}
}

func main() {
	fmt.Println("============================================================")
	fmt.Println("OFAuth Go SDK - Comprehensive Sandbox Test")
	fmt.Println("============================================================")
	fmt.Printf("\nBase URL: %s\n", baseURL)
	fmt.Printf("Connection ID: %s\n", connectionID)
	fmt.Printf("API Key: %s...\n", apiKey[:20])

	client := ofauth.NewClient(apiKey).
		WithBaseURL(baseURL).
		WithConnectionID(connectionID)

	ctx := context.Background()

	testAccountAPI(client, ctx)
	testSelfAPI(client, ctx)
	testPostsAPI(client, ctx)
	testMessagesAPI(client, ctx)
	testSubscribersAPI(client, ctx)
	testSubscriptionsAPI(client, ctx)
	testUsersAPI(client, ctx)
	testUserListsAPI(client, ctx)
	testVaultAPI(client, ctx)
	testEarningsAPI(client, ctx)
	testAnalyticsAPI(client, ctx)
	testPromotionsAPI(client, ctx)
	testMutationAPIs(client, ctx)

	total := results.passed + results.failed
	fmt.Println("\n============================================================")
	fmt.Printf("Results: %d/%d passed\n", results.passed, total)

	if results.failed > 0 {
		fmt.Println("\nFailed tests:")
		for _, e := range results.errors {
			fmt.Printf("  - %s: %s\n", e.test, e.error)
		}
		os.Exit(1)
	}
}
