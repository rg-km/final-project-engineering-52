import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  TableContainer,
} from '@chakra-ui/react'
import { useEffect } from 'react'
import { useComment } from '../../Database/useComment';

function Comment(){
  const {comments,fetch} = useComment(s => s) 
  useEffect(() => {
      fetch();
    }, []);
  return(
      <TableContainer>

<Table variant='simple'>
  <Thead>
    <Tr>
      <Th>ID</Th>
      <Th>Content</Th>
      <Th>Created_at</Th>
    </Tr>
  </Thead>
  <Tbody>
  {comments?.map((item, index) => (
    <Tr key={index}>
      <Td>{item.id}</Td>
      <Td>{item.content}</Td>
      <Td>{item.created_at}</Td>
    </Tr>
  ))}
  </Tbody>
  
</Table>
</TableContainer>
  )
}

export default Comment