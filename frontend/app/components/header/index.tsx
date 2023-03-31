import { Button, Flex, Heading, HStack, Show } from "@chakra-ui/react";
import { Link } from "@remix-run/react";
import { NotificationIcon } from "../icon/notification";

export const Header = ({ title, linkTo }: { title: string; linkTo?: string }) => {
  return (
    <HStack pt={10} mb={10} w={"full"} justifyContent={"space-between"}>
      <Heading color={"blackAlpha.800"} fontSize={"2xl"}>
        {title}
      </Heading>

      <HStack spacing={5}>
        {linkTo ? (
          <>
            <Show above="md">
              <Link to={linkTo}>
                <Button size={"sm"} colorScheme={"pink"} bgColor={"pink.400"}>
                  CREATE NEW
                </Button>
              </Link>
            </Show>
            <Show below="md">
              <Link to={linkTo}>
                <Button size={"sm"} colorScheme={"pink"} bgColor={"pink.400"}>
                  CREATE
                </Button>
              </Link>
            </Show>
          </>
        ) : null}

        <Flex justifyItems={"center"} alignItems={"center"}>
          <NotificationIcon boxSize={8} color={"blackAlpha.800"} />
        </Flex>
      </HStack>
    </HStack>
  );
};
