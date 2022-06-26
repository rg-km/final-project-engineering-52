import {
  Container,
  Stack,
  Flex,
  Heading,
  Text,
  Button,
  Image,
  Icon,

  // createIcon,
  useColorModeValue,
} from "@chakra-ui/react";

import { SimpleGrid } from "@chakra-ui/react";
import { IoAnalyticsSharp, IoLogoBitcoin, IoSearchSharp } from "react-icons/io5";
import { ReactElement } from "react";

function SplitWithImage() {
  return (
    <Container maxW={"full"} py={12} bg={useColorModeValue("pink.50", "pink.800")} mt={0} centerContent overflow="hidden">
      <SimpleGrid columns={{ base: 1, md: 2 }} spacing={10}>
        <Flex>
          <Image rounded={"md"} alt={"feature image"} src={"./about us.png"} objectFit={"cover"} />
        </Flex>
        <Stack spacing={5}>
          <Heading>GO-SCHOLARSHIP</Heading>
          <Text color={"black"} fontSize={"lg"}>
            go-scholarship adalah website yang membantu anda untuk mencari dan memfilter informasi beasiswa yang anda butuhkan. menggunakan website kami, anda dapat menggunakan fitur favorite pada postingan website sehingga anda dapat
            membaca ulang informasi beasiswa dengan mudah tanpa harus mencari ulang.
          </Text>
        </Stack>
      </SimpleGrid>
    </Container>
  );
}

export default function Hero() {
  return (
    <Container maxW={"full"} w={"100%"} mr={"0"} px={"0"}>
      <Stack px={"5"} align={"center"} spacing={{ base: 8, md: 10 }} py={{ base: 20, md: 28 }} direction={{ base: "column", md: "row" }}>
        <Stack flex={1} spacing={{ base: 5, md: 10 }}>
          <Heading lineHeight={1.1} fontWeight={600} fontSize={{ base: "3xl", sm: "4xl", lg: "6xl" }}>
            <br />
            <Text align={"left"} as={"span"} color={"black"}>
              Pelajari Segalanya dan Jelajahi Keterampilan Anda
            </Text>
          </Heading>
          <Text color={"grey"}>Pelajari Segalanya Dan Jelajahi Keterampilan Anda</Text>
          <Stack spacing={{ base: 4, sm: 6 }} direction={{ base: "column", sm: "row" }}>
            <Button rounded={"md"} size={"lg"} fontWeight={"normal"} px={6} colorScheme={"pink"} bg={"pink.400"} _hover={{ bg: "pink.500" }}>
              Mulai
            </Button>
          </Stack>
        </Stack>
        <Flex flex={1} justify={"center"} align={"center"} position={"relative"} w={"full"}>
          <Blob w={"150%"} h={"150%"} position={"absolute"} top={"-20%"} left={0} zIndex={-1} color={useColorModeValue("red.50", "red.400")} />
          <Image alt={"Hero Image"} fit={"cover"} align={"center"} w={"100%"} h={"100%"} src={"./home.png"} />
          {/* </Box> */}
        </Flex>
      </Stack>
      <SplitWithImage />
    </Container>
  );
}

export const Blob = (props) => {
  return (
    <Icon width={"100%"} viewBox="0 0 578 440" fill="none" xmlns="http://www.w3.org/2000/svg" {...props}>
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M239.184 439.443c-55.13-5.419-110.241-21.365-151.074-58.767C42.307 338.722-7.478 282.729.938 221.217c8.433-61.644 78.896-91.048 126.871-130.712 34.337-28.388 70.198-51.348 112.004-66.78C282.34 8.024 325.382-3.369 370.518.904c54.019 5.115 112.774 10.886 150.881 49.482 39.916 40.427 49.421 100.753 53.385 157.402 4.13 59.015 11.255 128.44-30.444 170.44-41.383 41.683-111.6 19.106-169.213 30.663-46.68 9.364-88.56 35.21-135.943 30.551z"
        fill="currentColor"
      />
    </Icon>
  );
};