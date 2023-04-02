import Segment from '../../Segment'
import { Box, Button } from '@chakra-ui/react'

const MyArticleCard = () => {
  return (
    <Segment>
      <Box>タイトル</Box>
      <Box>作成日時</Box>
      <Box>更新日時</Box>
      <Box display='flex' justifyContent='center' gap={2}>
        <Button>編集</Button>
        <Button>削除</Button>
      </Box>
    </Segment>
  )
}

export default MyArticleCard
