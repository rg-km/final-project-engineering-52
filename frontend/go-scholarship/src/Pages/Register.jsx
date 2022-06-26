import { EmailIcon } from "@chakra-ui/icons";
import {
  Flex,
  Box,
  // FormControl,
  // FormLabel,
  Input,
  Checkbox,
  Stack,
  Link,
  Button,
  Heading,
  Text,
  useColorModeValue,
  InputGroup,
  InputLeftElement,
  Icon,
} from "@chakra-ui/react";
import { useEffect, useRef, useState } from "react";
// import { PhoneIcon } from "@chakra-ui/icons";
// import { CgProfile } from "react-icons";
import { MdBook, MdCreditCard, MdPassword } from "react-icons/md";
import { useLocation, useNavigate } from "react-router-dom";
import { useAuth } from "../Database/useAuth";

export default function Register() {
  const location = useLocation();
  const asal = location?.state?.from?.pathname || "/"
  const navigate = useNavigate();
  const { register, user } = useAuth((state) => state);
  const [userForm, setUserForm] = useState({
    name: "",
    pendidikan: "",
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
      [e.target.name]: e.target.name == "image" ? e.target.files[0] : e.target.value,
    });
  }
  const handleRegister = (e) => {
    e.preventDefault()
    let fdata = new FormData()
    fdata.append("name", userForm.name)
    fdata.append("pendidikan", userForm.pendidikan)
    fdata.append("image", userForm.image)
    fdata.append("email", userForm.email)
    fdata.append("password", userForm.password)
    register(fdata, navigate);
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
          <Heading fontSize={"4xl"}>Daftar Akun</Heading>
          <Text fontSize={"lg"} color={"gray.600"}>
            Sudah Punya Akun Go-Scholarship ?{" "}
            <Link
            to="/login"
             color={"pink.400"}>Masuk</Link>
          </Text>
        </Stack>
        <Box
          rounded={"lg"}
          bg={useColorModeValue("white", "gray.700")}
          boxShadow={"lg"}
          p={8}
        >
          <Stack spacing={4}>
            <InputGroup>
              <InputLeftElement
                pointerEvents="none"
                children={<Icon as={MdCreditCard} color="pink.300" />}
              />
              <Input
                type="text"
                autoComplete="off"
                name="name"
                placeholder="Masukkan Nama Lengkap"
                onChange={changeHandler}
              />
            </InputGroup>

            <InputGroup>
              <InputLeftElement
                pointerEvents="none"
                children={<Icon as={MdBook} color="pink.300" />}
              />
              <Input
                type="text"
                autoComplete="off"
                name="pendidikan"
                placeholder="Masukkan Pendidikan"
                onChange={changeHandler}
              />
            </InputGroup>
            <InputGroup>
              <InputLeftElement
                pointerEvents="none"
                children={<Icon as={EmailIcon} color="pink.300" />}
              />
              <Input
                type="email"
                autoComplete="off"
                name="email"
                placeholder="Masukkan Email"
                onChange={changeHandler}
              />
            </InputGroup>
            <InputGroup>
              <InputLeftElement
                pointerEvents="none"
                children={<Icon as={EmailIcon} color="pink.300" />}
              />
              <Input
                type="file"
                name="image"
                autoComplete="off"
          
                onChange={changeHandler}
              />
            </InputGroup>
            <InputGroup>
              <InputLeftElement
                pointerEvents="none"
                children={<Icon as={MdPassword} color="pink.300" />}
              />
              <Input
                type="password"
                name="password"
                autoComplete="off"
                placeholder="Masukkan Password"
                onChange={changeHandler}
              />
            </InputGroup>
            <Stack spacing={10}>
              <Stack
                direction={{ base: "column", sm: "row" }}
                align={"start"}
                justify={"space-between"}
              >
                <Checkbox textAlign={"left"}>
                  Dengan mendaftar, kamu setuju untuk mengikuti{" "}
                  <Link color={"pink.400"}>Syarat Penggunaan</Link> dan{" "}
                  <Link color={"pink.400"}>Kebijakan Privasi.</Link>
                </Checkbox>
              </Stack>
              <Button
                bg={"pink.400"}
                onClick={handleRegister}
                color={"white"}
                _hover={{
                  bg: "pink.500",
                }}
              >
                Sign Up
              </Button>
            </Stack>
          </Stack>
        </Box>
      </Stack>
    </Flex>
  );
}
