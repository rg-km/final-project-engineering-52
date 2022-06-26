import {
    Heading,
    Avatar,
    Box,
    Center,
    Stack,
    VStack,
    Button,
    useColorModeValue,
    Stat,
    StatLabel,
    StatHelpText,
  } from '@chakra-ui/react';
import { useAuth } from '../Database/useAuth';
  
  export default function SocialProfileSimple() {
    const {user} = useAuth(s => s)
    
    return (
      <Center py={6}>
        <Box
          maxW={'320px'}
          w={'full'}
          bg={useColorModeValue('white', 'gray.900')}
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
          <Stat>
          <StatLabel>Nama Lengkap</StatLabel>
          <StatHelpText>{user?.name} </StatHelpText>
          </Stat>
          <Stat>
          <StatLabel>Pendidikan</StatLabel>
          <StatHelpText>sd</StatHelpText>
          </Stat>
          <Stat>
          <StatLabel>Email</StatLabel>
          <StatHelpText>{user?.email}</StatHelpText>
          </Stat>
          </VStack>
          <center mt={8} >
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
              Postingan Saya
            </Button>
            </center>
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