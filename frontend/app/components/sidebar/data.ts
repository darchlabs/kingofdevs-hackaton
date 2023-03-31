export const Items: {
  section: string;
  path: string;
  to: string;
  separator?: boolean;
}[] = [
  {
    section: "Metrics",
    path: "metrics",
    to: "/",
  },
  {
    section: "Smart Contracts",
    path: "smartcontracts",
    to: "/smartcontracts",
  },
  {
    section: "Users",
    path: "users",
    to: "/users",
    separator: true,
  },
  {
    section: "Settings",
    path: "settings",
    to: "/settings",
  },
];
