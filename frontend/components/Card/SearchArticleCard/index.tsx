import Segment from '../../Segment'
import { Box } from '@chakra-ui/react'

const SearchArticleCard = () => {
  return (
    <Segment>
      <Box>記事サムネイル画像</Box>
      <Box>タイトル</Box>
      <Box>作者</Box>
      <Box>作成日時</Box>
      <Box>更新日時</Box>
      <Box>価格</Box>
      <Box>販売数</Box>
    </Segment>
  )
}

export default SearchArticleCard
