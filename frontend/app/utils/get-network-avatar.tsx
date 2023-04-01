import type { Network } from "darchlabs";

import { PolygonAvatarIcon, EthereumAvatarIcon, AvalancheAvatarIcon, BaseAvatarIcon } from "../components/icon";

export const GetNetworkAvatar = (network: Network) => {
  switch (network) {
    case "polygon":
      return <PolygonAvatarIcon boxSize={12} />;
    case "ethereum":
      return <EthereumAvatarIcon boxSize={12} />;
    case "avalanche":
      return <AvalancheAvatarIcon boxSize={12} />;
  }

  return <BaseAvatarIcon boxSize={12} />;
};
