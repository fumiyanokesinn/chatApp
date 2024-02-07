import { Box } from "@yamada-ui/react";
import { Header } from "./Header";

export const Login = () => {
  return (
    <div>
      <div>
        <Header title="ログインページ" />
      </div>
      <div className="absolute inset-0 flex items-center justify-center">
        <Box p="md" rounded="md" bg="#d3d5da" color="black">
          ログイン ページ
        </Box>
      </div>
    </div>
  );
};
