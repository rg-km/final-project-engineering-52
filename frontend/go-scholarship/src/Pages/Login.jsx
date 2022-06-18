import {
  Flex,
  Box,
  FormControl,
  FormLabel,
  Input,
  Checkbox,
  Stack,
  // Link,
  Button,
  Heading,
  // Text,
  useColorModeValue,
} from '@chakra-ui/react';
import { useRef, useState } from 'react';
import { useAuth } from '../Database/useAuth';

export default function Login() {
  const {login}= useAuth(state =>state )
  const [userForm,setUserForm] = useState({
    email : '',
    password : ''
  })
  const ref = useRef()
  function changeHandler(e){
    setUserForm({
      ...userForm,
      [e.target.name]: e.target.value
    })
  }
  const handleLogin = () =>{
    login(userForm)
  } 
  return (
    <Flex
      minH={'100vh'}
      align={'center'}
      justify={'center'}
      bg={useColorModeValue('gray.50', 'gray.800')}>
      <Stack spacing={8} mx={'auto'} maxW={'lg'} py={12} px={6}>
        <Stack align={'center'}>
          <Heading fontSize={'4xl'}>GO - SCHOLARSHIP</Heading>
          
          {/* <Text fontSize={'lg'} color={'gray.600'}>
            to enjoy all of our cool <Link color={'blue.400'}>features</Link> ✌️
          </Text> */}
        </Stack>
        <Box
          rounded={'lg'}
          bg={useColorModeValue('white', 'gray.700')}
          boxShadow={'lg'}
          p={8}>
          <Stack spacing={4}>
            <FormControl id="email">
              <FormLabel>Email address</FormLabel>
              <Input type="email" name='email' onChange={changeHandler} ref={ref} />
            </FormControl>
            <FormControl id="password" >
              <FormLabel>Password</FormLabel>
              <Input type="password" onChange={changeHandler} name='password' ref={ref}/>
            </FormControl>
            
            <Stack spacing={10}>
              <Stack
                direction={{ base: 'column', sm: 'row' }}
                align={'start'}
                justify={'space-between'}>
                <Checkbox>Ingat saya</Checkbox>
                {/* <Link color={'blue.400'}>Forgot password?</Link> */}
              </Stack>
              <Button
                bg={'pink.400'}
                onClick={handleLogin}
                color={'white'}
                _hover={{
                  bg: 'pink.500',
                }}>
                Masuk
              </Button>

            </Stack>
          </Stack>
        </Box>
      </Stack>
    </Flex>
  );
}