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
  Spacer,
  useBoolean,
} from "@yamada-ui/react";
import { BuckButton } from "../components/BuckButton";
import { useState } from "react";
import useCreateAccount from "../hooks/useCreateAccount";

export const CreateAccount = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [show, { toggle }] = useBoolean();

  const { createAccount, message } = useCreateAccount();

  const onChangeName = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.currentTarget.value);
  };
  const onChangeEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.currentTarget.value);
  };
  const onChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.currentTarget.value);
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault(); // フォームのデフォルトの送信を防止
    createAccount({ name, email, password });
  };
  return (
    <div className="flex flex-col items-center justify-center m-4 sm:h-screen lg:py-0">
      <Box
        p="md"
        rounded="md"
        bg="#1c1e21"
        className="w-full md:mt-0 sm:max-w-md xl:p-0 text-white"
      >
        <Container>
          <Heading size="lg">Create your account</Heading>
          <form onSubmit={handleSubmit}>
            <FormControl label="Name">
              <Input placeholder="Alice" bg="#434851" onChange={onChangeName} />
            </FormControl>
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
          </form>
          {message}
          <Flex className="mt-4">
            <BuckButton />
            <Spacer />
            <Button colorScheme="success" type="submit">
              Register
            </Button>
          </Flex>
        </Container>
      </Box>
    </div>
  );
};
