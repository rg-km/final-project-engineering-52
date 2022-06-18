import {
  Box,
  Container,
  Heading,
  SimpleGrid,
  Text,
  Stack,
} from "@chakra-ui/react";

export default function GridListWithHeading() {
  return (
    <Box p={200}>
      <Stack spacing={4} as={Container} maxW={"6xl"} textAlign={"left"}>
        <Heading fontSize={"3xl"}>FAQ</Heading>
        <Text color={"black.600"} fontSize={"xl"}></Text>
      </Stack>
      <Container maxW={"6xl"} mt={10}>
        <SimpleGrid columns={{ base: 1 }} spacing={10} textAlign={"left"}>
          <Stack spacing={3}>
            <Heading fontSize={"2xl"}>
              <Text as={"span"} color={"pink.400"}>
                1. Apa itu Go Scholarship?
              </Text>
            </Heading>
            <Text color={"black.500"}>
              GO-SCHOLARSHIP adalah sebuah platform yang menyediakan informasi
              mengenai beasiswa.
            </Text>
            <Heading fontSize={"2xl"}>
              <Text as={"span"} color={"pink.400"}>
                2. Kenapa Harus memilih Go Scholarship ?
              </Text>
            </Heading>
            <Text color={"black.500"}>
              GO-SCHOLARSHIP memiliki Daftar Beasiswa yang bisa Anda pilih
              sesuai dengan kebutuhan Anda.
            </Text>
            <Heading fontSize={"2xl"}>
              <Text as={"span"} color={"pink.400"}>
                3. Siapa saja yang dapat mendaftar ?
              </Text>
            </Heading>
            <Text color={"black.500"}>
              Go Scholarship dapat di kunjungi oleh seluruh siswa dan mahasiswa
              yang berdomisili di Indonesia.
            </Text>
          </Stack>
        </SimpleGrid>
      </Container>
    </Box>
  );
}
