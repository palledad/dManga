import Segment from '../../components/Segment'
import Header from '../../components/Header'
import { Grid, GridItem, Input } from '@chakra-ui/react'
import SearchArticleCard from '../../components/Card/SearchArticleCard'
import Link from 'next/link'
import { useEffect } from 'react'
const MyPage = () => {
  useEffect(() => {
    console.log('useEffect')
    const articleAlias = ['string', '2', '3', 'eeee', 'aaa', 'onvm8flrpr']

    const headers = {
      accept: 'application/json',
    }

    const fetchArticle = async (alias: string) => {
      await fetch(`http://localhost:3000/api/articles/${alias}`, {
        method: 'GET',
        headers: headers,
      }).then((response) => {
        if (!response.ok) {
          console.log('error', response)
        } else {
          console.log('ok', response)
        }
      })
    }

    for (let i = 0; i < articleAlias.length; i++) {
      fetchArticle(articleAlias[i])
    }
  }, [])
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
