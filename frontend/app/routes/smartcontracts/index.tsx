import { HStack } from "@chakra-ui/react";
import { Outlet, useLoaderData } from "@remix-run/react";
import type { LoaderFunction, LoaderArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import type { ListEventsResponse } from "darchlabs";

import { Table } from "../../components/table";
import { Synchronizers } from "../../models/synchronizers.server";
import { EmptyTable, TableItem } from "../../components/smartcontracts";

type SynchronizersLoaderData = {
  events: ListEventsResponse;
};

export const loader: LoaderFunction = async ({ request }: LoaderArgs) => {
  const events = (await Synchronizers.listEvents()) as ListEventsResponse;

  // get sort query param
  const url = new URL(request.url);
  const sort = url.searchParams.get("sort") || "desc";

  // sort events
  if (sort === "asc") {
    events.data.sort((a, b) => {
      return new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime();
    });
  } else if (sort === "desc") {
    events.data.sort((a, b) => {
      return new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime();
    });
  }

  return json<SynchronizersLoaderData>({ events });
};

export function ErrorBoundary({ error }: { error: Error }) {
  console.error(error);
  return <>here in error section</>;
}

export default function App() {
  const {
    events: { data },
  } = useLoaderData<SynchronizersLoaderData>();

  return (
    <>
      <Outlet />

      <HStack justifyContent={"center"} w={"full"} pt={"20px"}>
        <Table
          title="smart contracts"
          columns={["Name", "network", "name", "status", "last updated", ""]}
          emptyTable={<EmptyTable createLink={"/smartcontracts/create"} />}
        >
          {data.map((item, index) => (
            <TableItem key={index} item={item} />
          ))}
        </Table>
      </HStack>
    </>
  );
}
