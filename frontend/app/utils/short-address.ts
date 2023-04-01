export const ShortAddress = (address: string, len: number = 5): string => {
  const start = address.substring(0, len + 2);
  const end = address.substring(address.length - len, address.length);

  return `${start}...${end}`;
};
