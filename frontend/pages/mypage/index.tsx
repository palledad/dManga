import Segment from '../../components/Segment'
import Header from '../../components/Header'
import {
  Box,
  Grid,
  GridItem,
  Tabs,
  TabList,
  TabPanels,
  Tab,
  TabPanel,
  Button,
  Input,
} from '@chakra-ui/react'

const MyPage = () => {
  return (
    <>
      <Header></Header>
      <Box height={'20em'} background={'gray.100'}></Box>
      <Segment>
        <Grid templateColumns='repeat(5, 1fr)' templateRows='repeat(2, 1fr)'>
          <GridItem colSpan={2}>ユーザー名</GridItem>
          <GridItem colSpan={3}>wallet address</GridItem>
          <GridItem colSpan={2}>自己紹介文</GridItem>
        </Grid>
      </Segment>
      <Segment>
        <Tabs variant='enclosed'>
          <TabList>
            <Tab>購入した記事</Tab>
            <Tab>自分の記事</Tab>
            <Tab>持っている収益NFT</Tab>
          </TabList>

          <TabPanels>
            <TabPanel>
              <Input placeholder='検索' background={'gray.100'} />
              <Segment>
                記事サムネイル画像
                <br />
                タイトル
                <br />
                作者
                <br />
                作成日時
                <br />
                更新日時
              </Segment>
            </TabPanel>
            <TabPanel>
              <Input placeholder='検索' background={'gray.100'} />
              <Segment>
                タイトル
                <br />
                作成日時
                <br />
                更新日時
                <br />
                <Button>編集</Button>
                <Button>削除</Button>
              </Segment>
            </TabPanel>
            <TabPanel>
              <Input placeholder='検索' background={'gray.100'} />
              <Segment>
                記事URL
                <br />
                記事タイトル
                <br />
                トークンID
              </Segment>
            </TabPanel>
          </TabPanels>
        </Tabs>
      </Segment>
    </>
  )
}

export default MyPage
