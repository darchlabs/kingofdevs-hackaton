import { useLocation, Link } from "@remix-run/react";
import { Box, Text, Image, VStack, HStack, Flex, useMediaQuery } from "@chakra-ui/react";

import { LogoIcon } from "../icon/logo";
import { LogoSquareIcon } from "../icon/logo-square";
import { getIconBySection } from "../../utils/get-icon-by-section";
import { Items } from "./data";

export function Sidebar() {
  const [AboveToLg] = useMediaQuery("(min-width: 62rem)");
  const { pathname } = useLocation();

  return (
    <VStack
      spacing={0}
      minW={{
        lg: "250px",
      }}
    >
      <Box pt={8} pb={8}>
        {AboveToLg ? <Image as={LogoIcon} boxSize={"135px"} /> : <Image as={LogoSquareIcon} boxSize={"55px"} />}
      </Box>

      <VStack width={"full"} alignItems={"stretch"} spacing={0}>
        {Items.map((item, index) => {
          const active = pathname === item.to;

          return (
            <Link key={index} to={item.to}>
              <HStack
                bg={active ? "pink.400" : "white"}
                h={14}
                color={active ? "white" : "blackAlpha.500"}
                _hover={{
                  backgroundColor: "pink.400",
                  color: "white",
                  cursor: "pointer",
                }}
                borderBottom={item.separator ? "1px" : "0px"}
                borderBottomColor={"blackAlpha.200"}
              >
                <Flex pl={6} pr={6} justifyContent={"center"} alignItems={"center"} alignContent="center">
                  {getIconBySection(item.path)}
                </Flex>
                {AboveToLg ? <Text fontSize={"md"}>{item.section}</Text> : null}
              </HStack>
            </Link>
          );
        })}
      </VStack>
    </VStack>
  );
}
