import { Box } from '@chakra-ui/react'
import Header from '../../components/Header'
import Segment from '../../components/Segment'

const MyAccount = () => {
  return (
    <>
      <Header></Header>
      <Segment>
        <Box>アカウント設定</Box>
        <Box>ウォレット連携</Box>
        <Box>アカウント名</Box>
      </Segment>
    </>
  )
}
export default MyAccount
