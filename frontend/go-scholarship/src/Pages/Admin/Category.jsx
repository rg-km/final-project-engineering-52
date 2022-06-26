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
import { useCategory } from '../../Database/useCategory';

function Category(){
  const {categories,fetch} = useCategory(s => s) 
  useEffect(() => {
      fetch();
    }, []);
  return(
      <TableContainer>

<Table variant='simple'>
  <Thead>
    <Tr>
      <Th>ID</Th>
      <Th>Category Name</Th>
      <Th>Created_at</Th>
    </Tr>
  </Thead>
  <Tbody>
  {categories?.map((item, index) => (
    <Tr key={index}>
      <Td>{item.id}</Td>
      <Td>{item.category_name}</Td>
      <Td>{item.created_at}</Td>
    </Tr>
  ))}
  </Tbody>
  
</Table>
</TableContainer>
  )
}

export default Category