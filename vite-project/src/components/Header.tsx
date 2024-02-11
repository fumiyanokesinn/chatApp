import { Box, Heading } from "@yamada-ui/react";

type Header = {
  title: string;
};

export const Header = (props: Header) => {
  return (
    <div>
      <Box p="md" bg="#2e3138">
        <div className="mx-40">
          <Heading size="md">{props.title}</Heading>
        </div>
      </Box>
    </div>
  );
};
