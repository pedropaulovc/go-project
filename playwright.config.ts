import { defineConfig } from "@playwright/test";

const PORT = process.env.E2E_PORT || "8080";

export default defineConfig({
  testDir: "./e2e",
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 1 : 0,
  workers: process.env.CI ? 1 : undefined,
  reporter: "html",
  timeout: process.env.CI ? 5_000 : 10_000,
  expect: { timeout: 2_000 },
  use: {
    baseURL: `http://localhost:${PORT}`,
    trace: "retain-on-failure",
  },
  webServer: {
    command: `go run .`,
    port: Number(PORT),
    reuseExistingServer: !process.env.CI,
  },
});
