import { Box, Button, Container, Flex, Heading, Input } from "@yamada-ui/react";
import { useState } from "react";

export const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const onChangeEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.currentTarget.value);
  };
  const onChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.currentTarget.value);
  };

  return (
    <div>
      <div className="flex flex-col items-center justify-center m-4 sm:h-screen lg:py-0">
        <Box
          p="md"
          rounded="md"
          bg="#1c1e21"
          className="md:mt-0 sm:max-w-md xl:p-0 text-white"
        >
          <Container>
            <Heading size="lg">Sign in to your account</Heading>
            <Input
              placeholder="メールアドレス"
              bg="#434851"
              onChange={onChangeEmail}
            />
            <Input
              placeholder="パスワード"
              bg="#434851"
              onChange={onChangePassword}
            />
            <Flex justify="center">
              <Button colorScheme="success">ログイン</Button>
            </Flex>
          </Container>
        </Box>
      </div>
    </div>
  );
};
