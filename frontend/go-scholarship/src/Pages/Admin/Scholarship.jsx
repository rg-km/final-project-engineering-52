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
import { useScholarship } from '../../Database/useScholarship';

function Scholarship(){
  const {scolarship,fetch} = useScholarship(s => s) 
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
  {scolarship?.map((item, index) => (
    <Tr key={index}>
      <Td>{item?.id}</Td>
      <Td>{item?.name}</Td>
      <Td>{item?.description}</Td>
      <Td>{item?.image}</Td>
      <Td>{item?.category?.category_name}</Td>
      <Td>{item?.users?.name}</Td>
    </Tr>
  ))}
  </Tbody>
  
</Table>
</TableContainer>
  )
}

export default Scholarship