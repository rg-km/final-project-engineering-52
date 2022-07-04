import {
    Table,
    Thead,
    Tbody,
    Tfoot,
    Tr,
    Th,
    Td,
    TableCaption,
    TableContainer,
  } from '@chakra-ui/react'
import { useEffect } from 'react'
import { useUsers } from '../../Database/useUsers'

function User(){
    const {users,fetch} = useUsers(s => s) 
    useEffect(() => {
        fetch();
      }, []);
    return(
        <TableContainer>

  <Table variant='simple'>
    <Thead>
      <Tr>
        <Th>ID</Th>
        <Th>Name</Th>
        <Th>Image</Th>
        <Th>Email</Th>
        <Th>Role</Th>
        <Th>Tgl Join</Th>
      </Tr>
    </Thead>
    <Tbody>
    {users?.map((item, index) => (
      <Tr key={index}>
        <Td>{item.id}</Td>
        <Td>{item.name}</Td>
        <Td>{item.image}</Td>
        <Td>{item.email}</Td>
        <Td>{item.role}</Td>
        <Td>{item.created_at}</Td>
      </Tr>
    ))}
    </Tbody>
    
  </Table>
</TableContainer>
    )
}

export default User