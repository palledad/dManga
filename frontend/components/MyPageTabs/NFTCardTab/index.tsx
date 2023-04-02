import { Grid, GridItem, Input } from '@chakra-ui/react'
import NFTCard from '../../Card/NFTCard'
import Link from 'next/link'

const NFTCardTab = () => {
  return (
    <>
      <Input placeholder='検索' background={'gray.100'} />
      <Grid templateColumns='repeat(3, 1fr)'>
        {[...Array(10)].map((x) => (
          <GridItem key={x} colSpan={1}>
            <Link href='articledetail'>
              <NFTCard />
            </Link>
          </GridItem>
        ))}
      </Grid>
    </>
  )
}
export default NFTCardTab
