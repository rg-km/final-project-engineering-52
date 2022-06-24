import {
  Flex,
  Box,
  FormControl,
  // FormLabel,
  Input,
  Checkbox,
  Stack,
  InputGroup,
  InputLeftElement,
  Icon,
  Link,
  Button,
  Heading,
  Text,
  useColorModeValue,
} from "@chakra-ui/react";
import { useRef, useState } from "react";
import { useAuth } from "../Database/useAuth";
import { MdOutlineEmail, MdPassword } from "react-icons/md";
import { useLocation, useNavigate } from "react-router-dom";
import { useEffect } from "react";

export default function Login() {
  const location = useLocation();
  const asal = location?.state?.from?.pathname || "/"
  const navigate = useNavigate();
  const { login, user } = useAuth((state) => state);
  const [userForm, setUserForm] = useState({
    email: "",
    password: "",
  });

  useEffect(() => {
    if(JSON.stringify(user) !== '{}') {
      navigate(asal, {
        replace: true
      })
    }
  }, [])

  const ref = useRef();
  function changeHandler(e) {
    setUserForm({
      ...userForm,
      [e.target.name]: e.target.value,
    });
  }
  const handleLogin = (e) => {
    e.preventDefault()
    login(userForm, navigate);
  };
  return (
    <Flex
      minH={"100vh"}
      align={"center"}
      justify={"center"}
      bg={useColorModeValue("gray.50", "gray.800")}
    >
      <Stack spacing={8} mx={"auto"} maxW={"lg"} py={12} px={6}>
        <Stack align={"center"}>
          <Heading fontSize={"4xl"}>GO - SCHOLARSHIP</Heading>

          {/* <Text fontSize={'lg'} color={'gray.600'}>
            to enjoy all of our cool <Link color={'blue.400'}>features</Link> ✌️
          </Text> */}
        </Stack>
        <Box
          rounded={"lg"}
          bg={useColorModeValue("white", "gray.700")}
          boxShadow={"lg"}
          p={8}
        >
          <Stack spacing={4}>
            <FormControl id="email">
              <InputGroup>
                <InputLeftElement
                  pointerEvents="none"
                  children={<Icon as={MdOutlineEmail} color="pink.300" />}
                />
                <Input
                  type="email"
                  name="email"
                  autoComplete="off"
                  placeholder="Masukkan Email"
                  onChange={changeHandler}
                  ref={ref}
                />
              </InputGroup>
            </FormControl>
            <FormControl id="password">
              <InputGroup>
                <InputLeftElement
                  pointerEvents="none"
                  children={<Icon as={MdPassword} color="pink.300" />}
                />
                <Input
                  type="password"
                  name="password"
                  placeholder="Masukkan Password"
                  onChange={changeHandler}
                  ref={ref}
                />
              </InputGroup>
            </FormControl>
            <Stack spacing={10}>
              <Stack
                direction={{ base: "column", sm: "row" }}
                align={"start"}
                justify={"space-between"}
              >
                <Checkbox>Ingat saya</Checkbox>
                {/* <Link color={'blue.400'}>Forgot password?</Link> */}
              </Stack>
              <Button
                bg={"pink.400"}
                onClick={handleLogin}
                color={"white"}
                _hover={{
                  bg: "pink.500",
                }}
              >
                Masuk
              </Button>
            </Stack>
            <Stack spacing={10}>
              <Text fontSize={"sm"} color={"black.600"}>
                Belum punya akun? <Link color={"pink.400"}>Daftar</Link>
              </Text>
            </Stack>
          </Stack>
        </Box>
      </Stack>
    </Flex>
  );
}
