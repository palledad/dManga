import Segment from '../../components/Segment'
import Header from '../../components/Header'
import { Grid, GridItem, Input } from '@chakra-ui/react'
import SearchArticleCard from '../../components/Card/SearchArticleCard'
import Link from 'next/link'

const MyPage = () => {
  return (
    <>
      <Header></Header>
      <Segment>
        <Input placeholder='検索' background={'gray.100'} />
        <Grid templateColumns='repeat(3, 1fr)'>
          {[...Array(10)].map((x) => (
            <GridItem key={x} colSpan={1}>
              <Link href='articleviewer'>
                <SearchArticleCard key={x} />
              </Link>
            </GridItem>
          ))}
        </Grid>
      </Segment>
    </>
  )
}

export default MyPage
