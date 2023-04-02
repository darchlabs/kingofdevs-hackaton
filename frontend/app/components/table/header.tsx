import { useLocation } from "react-router-dom";
import { HStack, VStack, Text, Button, Icon } from "@chakra-ui/react";
import { RiSortAsc, RiSortDesc } from "react-icons/ri";
import { Link } from "@remix-run/react";

export function Header({ title, length, subHeader }: { title: string; length: number; subHeader?: JSX.Element }) {
  // load hooks
  const { pathname, search } = useLocation();

  // get sort query param
  const queryParams = new URLSearchParams(search);
  const sort = queryParams.get("sort") || "desc";

  // set sort icon
  let sortIcon = RiSortDesc;
  if (sort === "asc") {
    sortIcon = RiSortAsc;
  }

  return (
    <HStack w={"full"} p={"32px"} justifyContent={"space-between"} alignItems="start">
      <VStack alignItems={"start"} spacing={1}>
        <Text fontWeight={"bold"} fontSize={"19px"}>
          All {title} ({length})
        </Text>
        {subHeader}
      </VStack>

      <HStack spacing={"24px"}>
        <HStack>
          <Link to={`${pathname}?sort=${sort === "asc" ? "desc" : "asc"}`}>
            <Button
              leftIcon={<Icon boxSize={5} color={"#C5C7CD"} as={sortIcon} />}
              colorScheme="pink"
              variant="ghost"
              color="#4B506D"
            >
              Sort
            </Button>
          </Link>
        </HStack>
      </HStack>
    </HStack>
  );
}
