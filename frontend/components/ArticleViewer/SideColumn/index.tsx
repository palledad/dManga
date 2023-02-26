import Segment from '../../Segment'
import { Box, Button } from '@chakra-ui/react'

const SideColumn = () => {
  return (
    <>
      <Segment>
        <Box>作者情報</Box>
        <Box>ユーザー名</Box>
        <Box>wallet address</Box>
        <Box>アイコン画像</Box>
        <Box>自己紹介</Box>
      </Segment>
      <Segment>
        <Box display='flex' gap={4} justifyContent='center'>
          <Button>購入</Button>
          <Button>ホルダーになる</Button>
        </Box>
      </Segment>
      <Segment>
        <Box>目次</Box>
      </Segment>
    </>
  )
}
export default SideColumn
