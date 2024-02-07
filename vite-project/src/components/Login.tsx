import { Box, Container, Heading, Input } from "@yamada-ui/react";
import { Header } from "./Header";

export const Login = () => {
  return (
    <div>
      <div>
        <Header title="ログインページ" />
      </div>
      <div className="absolute inset-0 flex items-center justify-center">
        <Box p="md" rounded="md" bg="#d3d5da" color="black">
          <Container>
            <Heading size="md">ログイン</Heading>
            <Input placeholder="メールアドレス" bg="white" />
            <Input placeholder="パスワード" bg="white" />
          </Container>
        </Box>
      </div>
    </div>
  );
};
