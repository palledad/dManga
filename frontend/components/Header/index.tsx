import { Image, GridItem, Button, Grid } from '@chakra-ui/react'
const Header = () => {
  return (
    <>
      <Grid templateColumns='repeat(5, 1fr)' background={'gray.500'} height={'3em'}>
        <GridItem colSpan={1}>
          <Image marginLeft={1} borderRadius='full' boxSize='30px' src='palledad_logo_1.jpg' />
        </GridItem>
        <GridItem colSpan={4} textAlign='right'>
          <Button size='sm' verticalAlign='middle' marginX={1}>
            記事検索
          </Button>
          <Button size='sm' marginX={1}>
            記事作成
          </Button>
          <Button size='sm' marginX={1}>
            アカウント
          </Button>
        </GridItem>
      </Grid>
    </>
  )
}

export default Header
