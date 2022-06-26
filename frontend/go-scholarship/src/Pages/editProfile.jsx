import {
    Heading,
    Avatar,
    Box,
    Center,
    Stack,
    VStack,
    Button,
    Input,
    FormControl,
    FormLabel,
  } from '@chakra-ui/react';
  
  export default function editProfile() {
    return (
      <Center py={6}>
        <Box
          maxW={'320px'}
          w={'full'}
          bg={('white', 'gray.900')}
          boxShadow={'2xl'}
          rounded={'lg'}
          p={6}
          textAlign={'center'}>
            <Heading lineHeight={1.1} fontSize={{ base: '2xl', sm: '3xl' }}>
           Profile
          </Heading>
          <Avatar
            size={'xl'}
            src={''}
            alt={'Avatar Alt'}
            mb={4}
            pos={'relative'}          
          />
          <VStack align={'left'} justify={'left'}  direction={'row'} mt={6}>
          <FormControl id="NamaLengkap" isRequired>
          <FormLabel>Nama Lengkap</FormLabel>
          <Input
            placeholder="NamaLengkap"
            _placeholder={{ color: 'gray.500' }}
            type="text"
          />
        </FormControl>
        <FormControl id="pendidikan" isRequired>
          <FormLabel>Pendidikan</FormLabel>
          <Input
            placeholder="Pendidikan"
            _placeholder={{ color: 'gray.500' }}
            type="text"
          />
          </FormControl>
        <FormControl id="email" isRequired>
          <FormLabel>Email address</FormLabel>
          <Input
            placeholder="your-email@example.com"
            _placeholder={{ color: 'gray.500' }}
            type="email"
          />
          </FormControl>
          </VStack>
          <Stack mt={8} direction={'row'} spacing={4}>
          <Button
              flex={1}
              fontSize={'sm'}
              rounded={'full'}
              bg={'pink.400'}
              color={'white'}
              boxShadow={
                '0px 1px 25px -5px rgb(66 153 225 / 48%), 0 10px 10px -5px rgb(66 153 225 / 43%)'
              }
              _hover={{
                bg: 'pink.400',
              }}
              _focus={{
                bg: 'pink.400',
              }}>
              Simpan
            </Button>
            <Button
              flex={1}
              fontSize={'sm'}
              rounded={'full'}
              bg={'pink.400'}
              color={'white'}
              boxShadow={
                '0px 1px 25px -5px rgb(66 153 225 / 48%), 0 10px 10px -5px rgb(66 153 225 / 43%)'
              }
              _hover={{
                bg: 'pink.400',
              }}
              _focus={{
                bg: 'pink.400',
              }}>
              Batal
            </Button>
          </Stack>
        </Box>
      </Center>
    );
  }