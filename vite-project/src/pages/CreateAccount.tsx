import { Box, Heading } from "@yamada-ui/react";
import { BuckButton } from "../components/BuckButton";

export const CreateAccount = () => {
  return (
    <div className="flex flex-col items-center justify-center m-4 sm:h-screen lg:py-0">
      <Box
        p="md"
        rounded="md"
        bg="#1c1e21"
        className="w-full md:mt-0 sm:max-w-md xl:p-0 text-white"
      >
        <Heading size="lg">Create your account</Heading>
        <BuckButton />
      </Box>
    </div>
  );
};
