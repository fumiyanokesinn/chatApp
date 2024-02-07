import { Box, Button, Container, Flex, Heading, Input } from "@yamada-ui/react";
import { Header } from "./Header";
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
      <div>
        <Header title="ログインページ" />
      </div>
      <div className="absolute inset-0 flex items-center justify-center">
        <Box p="md" rounded="md" bg="#d3d5da" color="black" className="w-1/4">
          <Container>
            <Heading size="lg">Third Chat</Heading>
            <Input
              placeholder="メールアドレス"
              bg="white"
              onChange={onChangeEmail}
            />
            <Input
              placeholder="パスワード"
              bg="white"
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
