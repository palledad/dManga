import { Box } from '@chakra-ui/react'
import Segment from '../Segment'

const ArticleColumn = () => {
  return (
    <>
      <Box height={'20em'} background={'gray.100'}>
        サムネイル画像
      </Box>
      <Segment>
        <Box>記事タイトル</Box>
      </Segment>
      <Segment>
        <Box>タグ一覧</Box>
      </Segment>
      <Segment>
        <Box>記事本文</Box>
      </Segment>
    </>
  )
}
export default ArticleColumn
