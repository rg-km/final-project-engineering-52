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
// import { PhoneIcon } from "@chakra-ui/icons";
// import { CgProfile } from "react-icons";
import { MdBook, MdCreditCard, MdPassword } from "react-icons/md";

export default function Login() {
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
            <Link color={"pink.400"}>Masuk</Link>
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
                placeholder="Masukkan Nama Lengkap"
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
                placeholder="Masukkan Pendidikan"
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
                placeholder="Masukkan Email"
              />
            </InputGroup>
            <InputGroup>
              <InputLeftElement
                pointerEvents="none"
                children={<Icon as={MdPassword} color="pink.300" />}
              />
              <Input
                type="password"
                autoComplete="off"
                placeholder="Masukkan Password"
              />
            </InputGroup>
            {/* <FormControl id="password">
              <FormLabel>Password</FormLabel>
              <Input type="password" />
            </FormControl> */}
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
                {/* <Link color={"blue.400"}>Forgot password?</Link> */}
              </Stack>
              <Button
                bg={"pink.400"}
                color={"white"}
                _hover={{
                  bg: "pink.500",
                }}
              >
                Sign in
              </Button>
            </Stack>
          </Stack>
        </Box>
      </Stack>
    </Flex>
  );
}
