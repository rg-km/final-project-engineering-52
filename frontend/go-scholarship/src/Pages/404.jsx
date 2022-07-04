import { Box, Heading, Text, Button } from '@chakra-ui/react';

export default function NotFound() {
  return (
    <Box  marginBottom={"14%"} textAlign="center" py={10} px={6}>
      <Heading
        display="inline-block"
        as="h2"
        size="2xl"
        backgroundColor={"pink.500"}
        // bgGradient="linear(to-r, teal.400, teal.600)"
        backgroundClip="text">
        404
      </Heading>
      <Text fontSize="18px" mt={3} mb={2}>
        Halaman Tidak Ditemukan
      </Text>
      <Text color={'black.500'} mb={6}>
         Silahkan kembali ke halaman sebelumnya
      </Text>

      <Button
        colorScheme="teal"
        backgroundColor={"pink.500"}
        color="white"
        variant="solid">
        Go to Home
      </Button>
    </Box>
  );
}