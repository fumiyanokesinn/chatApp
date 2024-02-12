import { Box, Heading } from "@yamada-ui/react";
import { ToggleDarkMode } from "./ToggleDarkMode";

type Header = {
  title: string;
};

export const Header = (props: Header) => {
  return (
    <div>
      <Box p="md" bg="#2e3138">
        <section className="sm:mx-40 flex justify-between items-center text-white">
          <Heading size="md">{props.title}</Heading>
          <ToggleDarkMode />
        </section>
      </Box>
    </div>
  );
};
