import { Heading, VStack, Text, Button } from "@chakra-ui/react";
import { Link } from "@remix-run/react";

export function EmptyTable({ createLink }: { createLink: string }) {
  return (
    <VStack
      w={"full"}
      maxW={"1000px"}
      alignItems={"center"}
      bg={"white"}
      borderWidth={"1px"}
      borderColor={"gray.300"}
      borderStyle={"solid"}
      borderRadius={"8px"}
      p={12}
      spacing={5}
      color={"blackAlpha.700"}
    >
      <Heading as="h2" size="lg" textAlign={"center"}>
        Get metrics from your smart contracts and their users
      </Heading>

      <Text fontSize="lg" textAlign={"center"}>
        DarchLabs offers several networks to execute smart contracts.
      </Text>

      <Text fontSize="md" textAlign={"center"}>
        With us, you can implement solutions that will allow you to obtain general metrics of your smart contracts and
        detailed analyses of the users who interact with them. This way, you can make more informed decisions and
        improve your results.
      </Text>

      <Link to={createLink}>
        <Button mt={7} size={"md"} colorScheme={"pink"} bgColor={"pink.400"}>
          CREATE SMART CONTRACT
        </Button>
      </Link>
    </VStack>
  );
}
