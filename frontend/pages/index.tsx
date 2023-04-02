import Header from '../components/Header'
import Segment from '../components/Segment'
import { Box } from '@chakra-ui/react'

export default function Home() {
  return (
    <>
      <Header></Header>
      <Segment>
        <Box>トップページ</Box>
      </Segment>
    </>
  )
}
