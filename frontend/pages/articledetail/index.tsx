import Header from '../../components/Header'
import Segment from '../../components/Segment'
import { Box } from '@chakra-ui/react'

const ArticleDetail = () => {
  return (
    <>
      <Header></Header>
      <Box height={'20em'} background={'gray.100'}>
        サムネイル画像
      </Box>
      <Segment>
        <Box>記事タイトル</Box>
      </Segment>
      <Segment>
        <Box>スペック</Box>
      </Segment>
      <Segment>
        <Box>タグ一覧</Box>
      </Segment>
    </>
  )
}
export default ArticleDetail
