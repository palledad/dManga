import { Grid, GridItem, Input } from '@chakra-ui/react'
import PurchasedArticleCard from '../../Card/PurchasedArticleCard'
import Link from 'next/link'

const PurchasedArticleTab = () => {
  return (
    <>
      <Input placeholder='検索' background={'gray.100'} />
      <Grid templateColumns='repeat(3, 1fr)'>
        {[...Array(10)].map((x) => (
          <GridItem key={x} colSpan={1}>
            <Link href='articleviewer'>
              <PurchasedArticleCard key={x} />
            </Link>
          </GridItem>
        ))}
      </Grid>
    </>
  )
}
export default PurchasedArticleTab
