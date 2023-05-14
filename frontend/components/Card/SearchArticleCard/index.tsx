import Segment from '../../Segment'
import { Box } from '@chakra-ui/react'
import { format, fromUnixTime } from 'date-fns'

type SearchArticleCardProps = {
  title: string
  author: string
  createdAt: number
  updatedAt: number
}
const SearchArticleCard = ({ title, author, createdAt, updatedAt }: SearchArticleCardProps) => {
  return (
    <Segment>
      <Box>記事サムネイル画像</Box>
      <Box>{title}</Box>
      <Box>{author}</Box>
      <Box>{format(fromUnixTime(createdAt), 'yyyy-MM-dd')}</Box>
      <Box>{format(fromUnixTime(updatedAt), 'yyyy-MM-dd')}</Box>
      {/*<Box>{price}</Box>*/}
      {/*<Box>{salesNum}</Box>*/}
    </Segment>
  )
}

export default SearchArticleCard
