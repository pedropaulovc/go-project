import { test, expect } from "@playwright/test";

test("home page shows hello world", async ({ page }) => {
  await page.goto("/");
  await expect(page.locator("h1")).toHaveText("Hello, World!");
});

test("health endpoint returns ok", async ({ request }) => {
  const response = await request.get("/health");
  expect(response.ok()).toBeTruthy();
  const body = await response.json();
  expect(body.status).toBe("ok");
});
