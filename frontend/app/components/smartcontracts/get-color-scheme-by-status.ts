import type { EventStatus } from "darchlabs";

export const GetColorSchemeByStatus = (status: EventStatus): string => {
  switch (status) {
    case "running":
      return "green";
    case "synching":
      return "yellow";
    case "error":
      return "red";
    case "stopped":
      return "gray";
  }
};
