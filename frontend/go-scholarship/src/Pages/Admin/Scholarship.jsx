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
  Button,
} from '@chakra-ui/react'
import { useEffect } from 'react'
import { useScholarship } from '../../Database/useScholarship';

function Scholarship(){
  const {scolarship,fetch,delete_data } = useScholarship(s => s) 
  useEffect(() => {
      fetch();
    }, []);
  return(
      <TableContainer>
        <Button>Add Scholarship</Button>
<Table variant='simple'>
  <Thead>
    <Tr>
      <Th>ID</Th>
      <Th>Name</Th>
      <Th>Image</Th>
      <Th>Description</Th>
      <Th>Category</Th>
      <Th>Author</Th>
      <Th>Action</Th>
    </Tr>
  </Thead>
  <Tbody>
  {scolarship?.map((item, index) => (
    <Tr key={index}>
      <Td>{item?.id}</Td>
      <Td>{item?.name}</Td>
      <Td>{item?.image}</Td>
      <Td>{item?.description}</Td>
      <Td>{item?.category?.category_name}</Td>
      <Td>{item?.user?.name}</Td>
      <Td><Button>Edit</Button>&nbsp;&nbsp;<Button onClick={() => delete(item?.id)}>Delete</Button></Td>
      
    </Tr>
  ))}
  </Tbody>
  
</Table>
</TableContainer>
  )
}

export default Scholarship