import {
    Heading,
    Avatar,
    Box,
    Center,
    Stack,
    VStack,
    Button,
    Stat,
    StatLabel,
    StatHelpText,
    Link,
  } from '@chakra-ui/react';
  
  export default function SocialProfileSimple() {
    return (
      <Center py={6}>
        <Box
          maxW={'320px'}
          w={'full'}
          bg={'gray.75'}
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
          <VStack justify={'left'}  direction={'row'} mt={6}>
          <Stat>
          <StatLabel>Nama Lengkap</StatLabel>
          <StatHelpText>Jane Doe</StatHelpText>
          </Stat>
          <Stat>
          <StatLabel>Pendidikan</StatLabel>
          <StatHelpText>Mahasiswa</StatHelpText>
          </Stat>
          <Stat>
          <StatLabel>Email</StatLabel>
          <StatHelpText>admin@gmail.com</StatHelpText>
          </Stat>
          </VStack>
          <Center>
          <Stack mt={15} direction={'row'} spacing={4}>
            <Link color='pink.500' href='./Listbea'>  
              Postingan Saya
              </Link>
              <Link color='pink.500' href='./editProfile'>
                Pengaturan
                </Link>
            </Stack>
            </Center>
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
              Hapus Akun
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
              Keluar
            </Button>
          </Stack>
        </Box>
      </Center>
    );
  }