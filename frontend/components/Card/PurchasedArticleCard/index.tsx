import Segment from '../../Segment'
import { Box } from '@chakra-ui/react'

const PurchasedArticleCard = () => {
  return (
    <Segment>
      <Box>記事サムネイル画像</Box>
      <Box>タイトル</Box>
      <Box>作者</Box>
      <Box>作成日時</Box>
      <Box>更新日時</Box>
    </Segment>
  )
}

export default PurchasedArticleCard
