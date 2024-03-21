import { Button } from "@yamada-ui/react";
import { Link } from "react-router-dom";

export const BuckButton = () => {
  return (
    <Link to="/">
      <Button colorScheme="primary">戻る</Button>
    </Link>
  );
};
