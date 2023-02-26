import { Grid, GridItem, Input } from '@chakra-ui/react'
import MyArticleCard from '../../Card/MyArticleCard'

const MyArticleCardTab = () => {
  return (
    <>
      <Input placeholder='検索' background={'gray.100'} />
      <Grid templateColumns='repeat(3, 1fr)'>
        {[...Array(10)].map((x) => (
          <GridItem key={x} colSpan={1}>
            <MyArticleCard />
          </GridItem>
        ))}
      </Grid>
    </>
  )
}
export default MyArticleCardTab
