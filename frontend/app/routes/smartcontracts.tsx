import { Outlet } from "@remix-run/react";
import { HStack, VStack } from "@chakra-ui/react";

import { Sidebar } from "../components/sidebar";
import { Header } from "../components/header";

export default function App() {
  return (
    <HStack alignItems={"start"} spacing={0}>
      <Sidebar />
      <VStack as={"section"} bg={"gray.50"} minW={0} w={"full"} h={"calc(100vh)"} pl={8} pr={8}>
        <Header title="Smart Contracts" linkTo="/smartcontracts/create" />
        <Outlet />
      </VStack>
    </HStack>
  );
}
