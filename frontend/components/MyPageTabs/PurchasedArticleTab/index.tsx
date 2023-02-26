import { Grid, GridItem, Input } from '@chakra-ui/react'
import PurchasedArticleCard from '../../Card/PurchasedArticleCard'

const PurchasedArticleTab = () => {
  return (
    <>
      <Input placeholder='検索' background={'gray.100'} />
      <Grid templateColumns='repeat(3, 1fr)'>
        {[...Array(10)].map((x) => (
          <GridItem key={x} colSpan={1}>
            <PurchasedArticleCard key={x} />
          </GridItem>
        ))}
      </Grid>
    </>
  )
}
export default PurchasedArticleTab
