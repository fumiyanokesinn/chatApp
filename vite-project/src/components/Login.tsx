import {
  Box,
  Button,
  Container,
  Flex,
  FormControl,
  Heading,
  Input,
  InputGroup,
  InputRightElement,
  useBoolean,
} from "@yamada-ui/react";
import { useState } from "react";
import useLogin from "../hooks/useLogin";

export const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [show, { toggle }] = useBoolean();

  const onChangeEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.currentTarget.value);
  };
  const onChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.currentTarget.value);
  };

  const { login, message } = useLogin();

  return (
    <div>
      <div className="flex flex-col items-center justify-center m-4 sm:h-screen lg:py-0">
        <Box
          p="md"
          rounded="md"
          bg="#1c1e21"
          className="w-full md:mt-0 sm:max-w-md xl:p-0 text-white"
        >
          <Container>
            <Heading size="lg">Sign in to your account</Heading>
            <FormControl label="Email address">
              <Input
                placeholder="example@company.com"
                bg="#434851"
                onChange={onChangeEmail}
              />
            </FormControl>
            <FormControl label="Password">
              <InputGroup>
                <Input
                  placeholder="••••••••"
                  bg="#434851"
                  onChange={onChangePassword}
                  type={show ? "text" : "password"}
                />
                <InputRightElement w="4.5rem" isClick>
                  <Button h="1.75rem" size="sm" onClick={toggle}>
                    {show ? "Hide" : "Show"}
                  </Button>
                </InputRightElement>
              </InputGroup>
            </FormControl>
            <Flex justify="center">
              <Button
                colorScheme="success"
                onClick={() => login({ email: email, password: password })}
              >
                ログイン
              </Button>
            </Flex>
            {message}
            <p className="text-sm font-light text-gray-400">
              Don’t have an account yet?{" "}
              <a href="#" className="font-medium hover:underline text-blue-400">
                Sign up
              </a>
            </p>
          </Container>
        </Box>
      </div>
    </div>
  );
};
