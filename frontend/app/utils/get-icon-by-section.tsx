import { VscOrganization, VscPieChart, VscSettingsGear } from "react-icons/vsc";
import { HiOutlineDocumentText } from "react-icons/hi";

export const GetIconBySection = (section: string) => {
  switch (section.toLowerCase()) {
    case "metrics":
      return <VscPieChart size={25} />;
    case "smartcontracts":
      return <HiOutlineDocumentText size={25} />;
    case "users":
      return <VscOrganization size={25} />;
    case "settings":
      return <VscSettingsGear size={25} />;
  }

  return null;
};
