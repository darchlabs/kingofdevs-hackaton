import { VStack } from "@chakra-ui/react";

import { Content } from "./content";
import { Header } from "./header";

export function Table({
  title,
  columns,
  subHeader,
  emptyTable,
  children,
}: {
  title: string;
  columns: string[];
  subHeader?: JSX.Element;
  emptyTable: JSX.Element;
  children: JSX.Element[];
}) {
  if (!children.length) {
    return emptyTable;
  }

  return (
    <VStack
      w={"full"}
      maxW={"1000px"}
      alignItems={"start"}
      bg={"white"}
      border={"1px solid #DFE0EB"}
      borderRadius={"8px"}
    >
      <Header title={title} length={children.length} subHeader={subHeader} />
      <Content columns={columns}>{children}</Content>
    </VStack>
  );
}
