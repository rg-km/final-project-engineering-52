import {
    Box,
    Button,
    FormControl,
    FormLabel,
    Heading,
    HStack,
    Input,
    Select,
    Stack,
  } from "@chakra-ui/react";
  
  export function CreateBeasiswa() {
    return (
      <>
      <Heading><h2>Create Beasiswa</h2></Heading>
        <Box
          marginTop={'100'}
          marginLeft={'auto'}
          marginRight={'auto'}
          boxSize={'lg'}
          boxShadow={"lg"}
          p={8}
        >
          <Stack spacing={4}>
            <HStack>
              <Box>
                <FormControl id="name" isRequired>
                  <FormLabel>Nama</FormLabel>
                  <Input type="text" />
                </FormControl>
              </Box>
              <Box>
                <FormControl id="kategori">
                  <FormLabel>kategori</FormLabel>
                  <Select placeholder='Pilih Kategori'>
                    <option value='option1'>Option 1</option>
                    <option value='option2'>Option 2</option>
                    <option value='option3'>Option 3</option>
                  </Select>
                </FormControl>
              </Box>
            </HStack>
            <FormLabel>Foto</FormLabel>
            <input type="file" />
            <FormControl id="deskripsi" isRequired>
              <FormLabel>Deskripsi</FormLabel>
              <Input type="teks" />
            </FormControl>
            <Stack spacing={10} pt={2}>
              <Button
                loadingText="Submitting"
                size="lg"
                bg={"pink.400"}
                color={"white"}
                _hover={{
                  bg: "blue.500",
                }}
              >
                Sign up
              </Button>
            </Stack>
          </Stack>
        </Box>
      </>
    );
  }
  