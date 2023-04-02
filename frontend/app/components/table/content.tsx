import {
  TableContainer,
  Table,
  Thead,
  Tr,
  Th,
  Tbody,
  TableCaption,
  HStack,
  Text,
  Stack,
  Select,
  Icon,
} from "@chakra-ui/react";
import { RiArrowLeftSLine, RiArrowRightSLine } from "react-icons/ri";

export function Content({ columns, children }: { columns: string[]; children: JSX.Element[] }) {
  return (
    <TableContainer w={"full"}>
      <Table variant="simple">
        <Thead>
          <Tr>
            {columns.map((column, index) => (
              <Th key={index} textTransform={"uppercase"}>
                {column}
              </Th>
            ))}
          </Tr>
        </Thead>
        <Tbody>{children}</Tbody>
        <TableCaption pt={0} mb={"8px"}>
          <HStack justifyContent={"end"} spacing={"15px"} mr={"24px"}>
            <HStack pr={"25px"}>
              <Text fontSize={"14px"} color={"#9FA2B4"}>
                Rows per page:
              </Text>
              <Stack spacing={3}>
                <Select variant="unstyled" size={"sm"} placeholder="5" color={"#4B506D"} />
              </Stack>
            </HStack>
            <HStack pr={"10px"}>
              <Text fontSize={"14px"} color={"#9FA2B4"}>
                {/* 1-{items.length} of {items.length} */}
              </Text>
            </HStack>
            <HStack>
              <Icon as={RiArrowLeftSLine} boxSize={5} color={"#9FA2B4"} />
            </HStack>
            <HStack>
              <Icon as={RiArrowRightSLine} boxSize={5} color={"#9FA2B4"} />
            </HStack>
          </HStack>
        </TableCaption>
      </Table>
    </TableContainer>
  );
}
